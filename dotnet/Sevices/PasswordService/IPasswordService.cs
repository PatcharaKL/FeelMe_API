namespace Project_FeelMe.Service.PassWordService;
 public partial interface IPassWordService
{
  Task<string> HashPassword(string password) ;
     Task<bool> CheckPassword(string password,string hashPassword);
}

