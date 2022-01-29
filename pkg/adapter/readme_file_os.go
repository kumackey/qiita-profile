package adapter

import (
	"bufio"
	"os"

	"github.com/kumackey/profile-updater/pkg/domain"
)

const filenameReadMe = "README.md"

type ReadmeFileOS struct{}

func (s ReadmeFileOS) Scan() (*domain.Profile, error) {
	f, err := os.Open(filenameReadMe)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []domain.Line
	for scanner.Scan() {
		lines = append(lines, domain.Line(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return &domain.Profile{
		Content: lines,
	}, nil
}

func (s ReadmeFileOS) Write(readme *domain.Profile) error {
	f, err := os.Create(filenameReadMe)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range readme.Content {
		_, err := f.WriteString(string(line) + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}
