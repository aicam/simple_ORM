package server

func (s *Server) Routes() {
	s.Router.GET("/check_token/:token", s.CheckToken())
	s.Router.POST("/admin_login", s.AdminLogin())
	s.Router.GET("/add_admin/:user/:pass", s.AddAdmin())
	s.Router.POST("/add_customer", s.CheckAuthentication(), s.AddCustomer())
	s.Router.GET("/get_all_customers", s.CheckAuthentication(), s.GetCustomers())
	s.Router.POST("/add_payment", s.CheckAuthentication(), s.AddPayment())
	s.Router.GET("/get_all_payments", s.CheckAuthentication(), s.GetPayments())
	s.Router.POST("/add_product", s.CheckAuthentication(), s.AddProduct())
	s.Router.GET("/get_all_products", s.CheckAuthentication(), s.GetProducts())
	s.Router.POST("/add_order/", s.CheckAuthentication(), s.AddOrder())
	s.Router.GET("/get_all_orders", s.CheckAuthentication(), s.GetOrders())
	s.Router.POST("/update_order", s.CheckAuthentication(), s.UpdateOrder())
	s.Router.POST("/remove_order", s.CheckAuthentication(), s.RemoveOrder())
}
