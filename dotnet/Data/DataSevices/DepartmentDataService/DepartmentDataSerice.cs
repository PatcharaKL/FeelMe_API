using Microsoft.EntityFrameworkCore;
using Project_FeelMe.Data;
using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.DepartmentDataService
{
    public class DepartmentDataSerice:IDepartmentDataSerice
    {
        private readonly FeelMeContext _dbContract;

        public DepartmentDataSerice(FeelMeContext dbContract)
        {
            _dbContract = dbContract;
        }
        public async Task<List<Department>> GetAllDepartmentAsync()
        {
            return await _dbContract.Departments.ToListAsync();
        }

        public async Task<Department> GetDepartmentByIdAsync(int id)
        {
           return await _dbContract.Departments.FirstOrDefaultAsync(department => department.DepartmentId == id);
        }

        public async Task InsertDepartmentAsync(Department department)
        {
            _dbContract.Add(department);
            await _dbContract.SaveChangesAsync();
        }

        public async Task InsertDepartmentAsync(List<Department> department)
        {
           _dbContract.AddRange(department);
            await _dbContract.SaveChangesAsync();
        }

        public async Task UpdateDepartmentAsync(List<Department> department)
        {
            _dbContract.UpdateRange(department);
            await _dbContract.SaveChangesAsync();
        }

        public async Task UpdateDepartmentsAsync(Department department)
        {
            _dbContract.Update(department);
            await _dbContract.SaveChangesAsync();
        }
    }
}