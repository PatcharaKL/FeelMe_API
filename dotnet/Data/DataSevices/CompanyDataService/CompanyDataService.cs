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
        public async Task<List<Company>> GetAllCompanyAsync()
        {
            return await _dbContract.Companies.ToListAsync();
        }

        public async Task<Company> GetCompanyByIdAsync(int id)
        {
           return await _dbContract.Companies.FirstOrDefaultAsync(Company => Company.CompanyId == id);
        }

        public async Task InsertCompanyAsync(Company Company)
        {
            _dbContract.Add(Company);
            await _dbContract.SaveChangesAsync();
        }

        public async Task InsertCompanyAsync(List<Company> Company)
        {
           _dbContract.AddRange(Company);
            await _dbContract.SaveChangesAsync();
        }

        public async Task UpdateCompanyAsync(List<Company> Company)
        {
            _dbContract.UpdateRange(Company);
            await _dbContract.SaveChangesAsync();
        }

        public async Task UpdateCompanysAsync(Company Company)
        {
            _dbContract.Update(Company);
            await _dbContract.SaveChangesAsync();
        }
    }
}