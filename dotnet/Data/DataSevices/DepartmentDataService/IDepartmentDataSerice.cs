using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.DepartmentDataService
{
    public interface IDepartmentDataSerice
    {
         Task<List<Department>> GetAllDepartmentAsync();
         Task<Department> GetDepartmentByIdAsync(int id);
         Task InsertDepartmentAsync(Department department);
         Task InsertDepartmentAsync(List<Department> department);
         Task UpdateDepartmentsAsync(Department department);
         Task UpdateDepartmentAsync(List<Department> department);
    }
}