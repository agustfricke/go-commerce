package models

type User struct {
  ID     int64
	Name          string  
	Email         string 
	Password      string
	Verified      bool 
	SocialID      string  
	Avatar        string 
	Otp_enabled   bool  
	Otp_verified  bool 

	Otp_secret    string
	Otp_auth_url  string
}

type SignUpInput struct {
    Name  string
    Email  string
    Password string
}
