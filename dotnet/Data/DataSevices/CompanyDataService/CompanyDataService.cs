using Microsoft.EntityFrameworkCore;
using Project_FeelMe.Data;
using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.CompanyDataService
{
    public class CompanyDataService:ICompanyDataService
    {
        private readonly FeelMeContext _dbContract;

        public CompanyDataService(FeelMeContext dbContract)
        {
            _dbContract = dbContract;
        }
        public virtual async Task<List<Company>> GetAllCompanyAsync()
        {
            return await _dbContract.Companies.ToListAsync();
        }

        public virtual async Task<Company> GetCompanyByIdAsync(int id)
        {
           return await _dbContract.Companies.FirstOrDefaultAsync(Company => Company.CompanyId == id);
        }

        public virtual async Task InsertCompanyAsync(Company Company)
        {
            _dbContract.Add(Company);
            await _dbContract.SaveChangesAsync();
        }

        public virtual async Task InsertListCompanyAsync(List<Company> Company)
        {
           _dbContract.AddRange(Company);
            await _dbContract.SaveChangesAsync();
        }

        public virtual async Task UpdateListCompanyAsync(List<Company> Company)
        {
            _dbContract.UpdateRange(Company);
            await _dbContract.SaveChangesAsync();
        }

        public virtual async Task UpdateCompanysAsync(Company Company)
        {
            _dbContract.Update(Company);
            await _dbContract.SaveChangesAsync();
        }
    }
}