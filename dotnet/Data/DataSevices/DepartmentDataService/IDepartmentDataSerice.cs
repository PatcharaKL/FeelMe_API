using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.DepartmentDataService
{
    public interface IDepartmentDataSerice
    {
         Task<List<Department>> GetAllDepartmentAsync();
         Task<Department> GetDepartmentByIdAsync(int id);
         Task InsertDepartmentAsync(Department department);
         Task InsertListDepartmentAsync(List<Department> department);
         Task UpdateDepartmentsAsync(Department department);
         Task UpdateListDepartmentAsync(List<Department> department);
    }
}