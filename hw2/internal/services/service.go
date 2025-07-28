package service

import (
	"bufio"
	"os"
)

type Service struct {
	prod Producer
	pres Presenter
}

func New(p Producer, pr Presenter) *Service {
	return &Service{
		prod: p,
		pres: pr,
	}
}

func (s *Service) mask(data string) string {
	if len(data) > 4 {
		return data[:4] + "****"
	}
	return "****"
}

// Главный метод запуска
func (s *Service) Run() error {
	input, err := s.prod.Produce()
	if err != nil {
		return err
	}

	var output []string
	for _, v := range input {
		output = append(output, s.mask(v))
	}

	if err := s.pres.Present(output); err != nil {
		return err
	}

	return nil
}

type FileProducer struct {
	path string
}

func NewFileProducer(path string) *FileProducer {
	return &FileProducer{path: path}
}

type Producer interface {
	Produce() ([]string, error)
}

func (f *FileProducer) Produce() ([]string, error) {
	file, err := os.Open(f.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

type Presenter interface {
	Present([]string) error
}

type FilePresenter struct {
	path string
}

func NewFilePresenter(path string) *FilePresenter {
	return &FilePresenter{path: path}
}

func (f *FilePresenter) Present(data []string) error {
	file, err := os.Create(f.path)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range data {
		if _, err := file.WriteString(line + "\n"); err != nil {
			return err
		}
	}
	return nil
}
