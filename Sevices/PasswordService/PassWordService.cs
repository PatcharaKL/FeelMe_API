namespace Project_FeelMe.Service.PassWordService;

    public partial class PassWordService:IPassWordService
{
     public virtual async Task<string> HashPassword(string password) 
        { 
            return await Task.FromResult<string>(BCrypt.Net.BCrypt.HashPassword(password));
        }
        public virtual async Task<bool> CheckPassword(string password,string hashPassword)
        {
            return await Task.FromResult<bool>(BCrypt.Net.BCrypt.Verify(password,hashPassword));
        }
}

