namespace dotnet.ViewModel
{
    public class UserConstants
    {
        public static List<AccountViewModels> Users = new List<AccountViewModels>()
        {
            new AccountViewModels() {  Email = "Eart.admin@email.com", Password = "MyPass_w0rd", Name = "Eart", Surname = "Kleap", PositionId = 1},
            new AccountViewModels() {  Email = "Sayfar@email.com", Password = "MyPass_w0rd", Name = "Sayfar", Surname = "Hongsaeng", PositionId = 2 },
        };
    }
}