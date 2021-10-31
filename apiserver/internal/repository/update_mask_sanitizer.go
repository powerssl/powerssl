package repository

type UpdateMaskSanitizer struct {
	allowed  map[string]struct{}
	internal map[string]struct{}
}

func NewUpdateMaskSanitizer() *UpdateMaskSanitizer {
	return &UpdateMaskSanitizer{
		allowed:  make(map[string]struct{}),
		internal: make(map[string]struct{}),
	}
}

func (s *UpdateMaskSanitizer) Allowed(paths ...string) *UpdateMaskSanitizer {
	for _, path := range paths {
		s.allowed[path] = struct{}{}
	}
	return s
}

func (s *UpdateMaskSanitizer) Internal(paths ...string) *UpdateMaskSanitizer {
	for _, path := range paths {
		s.internal[path] = struct{}{}
	}
	return s
}

func (s *UpdateMaskSanitizer) Sanitize(paths []string, internal bool) []string {
	n := 0
	for _, path := range paths {
		if _, ok := s.allowed[path]; ok {
			paths[n] = path
			n++
		}
	}
	allowedPaths := paths[:n]
	var internalPaths []string
	if internal {
		n = 0
		for _, path := range paths {
			if _, ok := s.internal[path]; ok {
				paths[n] = path
				n++
			}
		}
		internalPaths = paths[:n]
	}
	return append(allowedPaths, internalPaths...)
}
