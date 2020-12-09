package server

func (s *Server) Routes() {
	s.Router.POST("/admin_login", s.AdminLogin())
	s.Router.GET("/add_admin/:user/:pass", s.CheckAuthentication(), s.AddAdmin())
	s.Router.POST("/add_customer", s.CheckAuthentication(), s.AddCustomer())
	s.Router.GET("/get_all_customers", s.CheckAuthentication(), s.GetCustomers())
}
