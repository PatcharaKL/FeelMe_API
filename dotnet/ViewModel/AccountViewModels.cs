namespace dotnet.ViewModel
{
    public class AccountViewModels
    {
        public AccountViewModels()
        {
            userName = new  UserName();
        }
        public string Email { get; set; } = null!;
        public int AccountId { get; set; } 
        public string Name { get; set; } = null!;
        public string Surname { get; set; } = null!;
         public string Password { get; set; } = null!;
         public int DepartmentId { get; set; }
        public int PositionId { get; set; }
        public int CompanyId { get; set; }
        public UserName userName  { get; set; }

        public class UserName  
        {
             public string userName { get; set; } 
        }
    }
}