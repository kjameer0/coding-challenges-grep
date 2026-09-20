package fileparse

import (
	"errors"
	"io/fs"
	"os"
	"slices"
)

//we want to return a file set and then try to remove and include files

var (
	ErrNoPathsProvided = errors.New("No path provided as universal set")
)

type ParseCfg struct {
	UniversePaths       []string
	IncludeFiles        []string
	ExcludeFiles        []string
	ExcludeDirs         []string
	outputFileSet       []string
	seenPaths           map[string]bool
	RecursiveResolution bool
}

type ParseOption func(*ParseCfg)

func WithIncludeFiles(files []string) ParseOption {
	return func(p *ParseCfg) {
		p.IncludeFiles = files
	}
}

func WithRecursiveResolution(on bool) ParseOption {
	return func(p *ParseCfg) {
		p.RecursiveResolution = on
	}
}
func WithExcludeFiles(files []string) ParseOption {
	return func(p *ParseCfg) {
		p.ExcludeFiles = files
	}
}
func WithExcludeDirs(dirs []string) ParseOption {
	return func(p *ParseCfg) {
		p.ExcludeDirs = dirs
	}
}
func WithUniversePaths(paths []string) ParseOption {
	return func(p *ParseCfg) {
		p.UniversePaths = paths
	}
}

func (p *ParseCfg) CustomWalkDirFuncGenerator() fs.WalkDirFunc {
	return func(path string, dir fs.DirEntry, err error) error {
		if dir.IsDir() && p.RecursiveResolution == false && path != "." {
			return fs.SkipDir
		}
		if dir.IsDir() {
			return nil
		}
		p.outputFileSet = append(p.outputFileSet, path)
		p.seenPaths[path] = true
		return nil
	}
}

func NewParseCfg(opts ...ParseOption) (*ParseCfg, error) {
	seenPaths := make(map[string]bool, 1000)
	c := &ParseCfg{
		UniversePaths: []string{},
		IncludeFiles:  []string{},
		ExcludeFiles:  []string{},
		ExcludeDirs:   []string{},
		outputFileSet: []string{},
		seenPaths:     seenPaths,
	}
	for _, opt := range opts {
		opt(c)
	}
	if len(c.UniversePaths) == 0 {
		return nil, ErrNoPathsProvided
	}
	return c, nil
}

func (p *ParseCfg) ConstructFileSet() ([]string, error) {
	//start from the universe and narrow by includes -> excludes files -> excludeDirs
	for _, path := range p.UniversePaths {
		if p.seenPaths[path] {
			continue
		}

		f, err := os.Stat(path)
		if err != nil {
			if err == fs.ErrNotExist {
				//TODO: debug log the lack of existence of a file
				continue
			}
			return nil, err
		}

		if !f.IsDir() {
			p.outputFileSet = append(p.outputFileSet, path)
			p.seenPaths[path] = true
			continue
		}
		//TODO: handle skipping a dir when recursion is active
		fs.WalkDir(os.DirFS(path), ".", p.CustomWalkDirFuncGenerator())
		p.seenPaths[path] = true
		// root of recursion
	}
	return slices.Clone(p.outputFileSet), nil
}
