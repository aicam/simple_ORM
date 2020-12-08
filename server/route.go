package server

func (s *Server) Routes() {
	s.Router.GET("/add_user/:username/:score", s.addUser())
}

