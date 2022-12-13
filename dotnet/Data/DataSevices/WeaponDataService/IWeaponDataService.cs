using Project_FeelMe.Models;
namespace dotnet.Data.DataSevices.WeaponDataService
{
    public interface IWeaponDataService
    {
         Task<List<Weapon>> GetAllWeaponAsync();
         Task<Weapon> GetWeaponByWeaponIdAsync(int id);
         Task UpdateWeaponAsync(Weapon weapon);
         Task InsertAsyncWeapon(Weapon weapon);
         Task UpdateWeaponAsync(List<Weapon> weapon);
    }
}