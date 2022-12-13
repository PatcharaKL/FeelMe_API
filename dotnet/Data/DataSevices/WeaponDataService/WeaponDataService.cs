using Microsoft.EntityFrameworkCore;
using Project_FeelMe.Data;
using Project_FeelMe.Models;
namespace dotnet.Data.DataSevices.WeaponDataService
{
    public class WeaponDataService:IWeaponDataService
    {
          private readonly FeelMeContext _dbContract;

            public WeaponDataService(FeelMeContext dbContract)
            {
                 _dbContract = dbContract;
            }
            public virtual async Task<List<Weapon>> GetAllWeaponAsync()
            {
                var data = await _dbContract.Weapons.ToListAsync();
                return data;
            }
            public  virtual async Task<Weapon> GetWeaponByWeaponIdAsync(int id)
            {
                return await _dbContract.Weapons.FirstOrDefaultAsync(weapon => weapon.WeaponsId == id);
            }
             public virtual async Task UpdateWeaponAsync(Weapon weapon)
             {
                 _dbContract.Update(weapon);
                await _dbContract.SaveChangesAsync();
             }
             public virtual async Task UpdateWeaponAsync(List<Weapon> weapon)
             {
                 _dbContract.UpdateRange(weapon);
                await _dbContract.SaveChangesAsync();
             }   
             public virtual async Task InsertAsyncWeapon(Weapon weapon)
             {
                _dbContract.Add(weapon);
                await _dbContract.SaveChangesAsync();
             }
    }
}