using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.CompanyDataService
{
    public interface ICompanyDataService
    {
         Task<List<Company>> GetAllCompanyAsync();
         Task<Company> GetCompanyByIdAsync(int id);
         Task InsertCompanyAsync(Company Company);
         Task InsertListCompanyAsync(List<Company> Company);
         Task UpdateCompanysAsync(Company Company);
         Task UpdateListCompanyAsync(List<Company> Company);
    }
}