public class PassWordService
{
     public string HashPassword(string password) 
        { 
            return BCrypt.Net.BCrypt.HashPassword(password);
        }
        public bool CheckPassword(string password,string hashPassword)
        {
            return BCrypt.Net.BCrypt.Verify(password,hashPassword) ;
        }
}