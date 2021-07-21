package server

func (s *Server) setRoutes() {
	s.router.HandleFunc("/metadata", s.handleMetadata())
}
