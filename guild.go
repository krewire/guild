// Package guild ships the Krewire Guild: a reusable AI agent setup
// (AGENTS.md, opencode.json, and a .agents/ preset) installable into any
// software project.
//
// The template is embedded and exposed as Template. Install copies it into a
// target directory, refusing to overwrite managed files unless forced.
package guild

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
)

// Template embeds the installable Guild template tree.
//
//go:embed all:template
var Template embed.FS

// managed is the canonical set of top-level paths Install writes into a
// target directory. Anything copied lives under one of these prefixes.
func managed() []string {
	return []string{
		"AGENTS.md",
		"opencode.json",
		".agents/agents",
		".agents/commands",
		".agents/skills",
		".agents/context",
		".agents/README.md",
	}
}

var (
	// ErrTargetMissing is returned when the target directory does not exist
	// or is empty.
	ErrTargetMissing = errors.New("install: target directory does not exist")
	// ErrConflicts is returned when a managed file already exists in the
	// target and the install is not forced.
	ErrConflicts = errors.New("install: managed files already exist")
)

// Option configures an Install call.
type Option func(*options)

type options struct {
	force  bool
	dryRun bool
}

// WithForce permits overwriting existing managed files in the target.
func WithForce() Option {
	return func(o *options) { o.force = true }
}

// WithDryRun validates and reports the would-be writes without touching the
// target directory.
func WithDryRun() Option {
	return func(o *options) { o.dryRun = true }
}

// Install copies the Guild template into target. It returns the destination
// paths that were (or, with WithDryRun, would be) written, sorted, and a
// typed error when the target is unusable or existing files conflict.
func Install(target string, opts ...Option) ([]string, error) {
	target = filepath.Clean(target)

	info, err := os.Stat(target)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, fmt.Errorf("%w: %s", ErrTargetMissing, target)
		}
		return nil, fmt.Errorf("install: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("%w: %s", ErrTargetMissing, target)
	}

	cfg := &options{}
	for _, apply := range opts {
		apply(cfg)
	}

	conflicts, err := detectConflicts(target)
	if err != nil {
		return nil, err
	}
	if len(conflicts) > 0 && !cfg.force && !cfg.dryRun {
		return nil, fmt.Errorf("%w: %s", ErrConflicts, listConflicts(target, conflicts))
	}

	relPaths, err := templatePaths()
	if err != nil {
		return nil, err
	}

	created := make([]string, 0, len(relPaths))
	for _, rel := range relPaths {
		dest := filepath.Join(target, rel)
		if cfg.dryRun {
			created = append(created, dest)
			continue
		}
		data, err := fs.ReadFile(Template, "template/"+rel)
		if err != nil {
			return nil, fmt.Errorf("install: read %s: %w", rel, err)
		}
		if err := writeDest(dest, data); err != nil {
			return nil, err
		}
		created = append(created, dest)
	}
	sort.Strings(created)
	return created, nil
}

// Managed returns the canonical top-level paths the installer manages, as
// relative paths inside target directories.
func Managed() []string {
	return append([]string(nil), managed()...)
}

// detectConflicts returns managed paths that already exist under target.
func detectConflicts(target string) ([]string, error) {
	var conflicts []string
	for _, m := range managed() {
		exists, err := anyExist(filepath.Join(target, m))
		if err != nil {
			return nil, err
		}
		if exists {
			conflicts = append(conflicts, m)
		}
	}
	sort.Strings(conflicts)
	return conflicts, nil
}

// anyExist reports whether any file or directory exists at p.
func anyExist(p string) (bool, error) {
	if _, err := os.Lstat(p); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, fmt.Errorf("install: stat %s: %w", p, err)
	}
	return true, nil
}

// templatePaths returns every file path inside the embedded template tree,
// relative to the template root.
func templatePaths() ([]string, error) {
	var rels []string
	err := fs.WalkDir(Template, "template", func(p string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		rel, err := relFor(p)
		if err != nil {
			return err
		}
		rels = append(rels, rel)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("install: walk template: %w", err)
	}
	sort.Strings(rels)
	return rels, nil
}

// relFor strips the leading "template/" prefix from an embedded path.
func relFor(p string) (string, error) {
	const prefix = "template/"
	if !hasPrefix(p, prefix) {
		return "", fmt.Errorf("install: unexpected embedded path %q", p)
	}
	return path.Clean(string(p[len(prefix):])), nil
}

func hasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// writeDest writes data to dest, creating parent directories as needed.
func writeDest(dest string, data []byte) error {
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return fmt.Errorf("install: mkdir %s: %w", filepath.Dir(dest), err)
	}
	if err := os.WriteFile(dest, data, 0o644); err != nil {
		return fmt.Errorf("install: write %s: %w", dest, err)
	}
	return nil
}

// listConflicts formats managed paths that already exist under target for
// error messages.
func listConflicts(target string, paths []string) string {
	out := make([]string, 0, len(paths))
	for _, p := range paths {
		out = append(out, filepath.Join(target, p))
	}
	return joinPaths(out)
}

func joinPaths(parts []string) string {
	result := ""
	for i, p := range parts {
		if i > 0 {
			result += ", "
		}
		result += p
	}
	return result
}
