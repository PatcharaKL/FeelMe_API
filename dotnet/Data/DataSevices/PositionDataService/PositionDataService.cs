using Microsoft.EntityFrameworkCore;
using Project_FeelMe.Data;
using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.PositionDataService
{
    public class PositionDataService:IPositionDataService
    {
        private readonly FeelMeContext _dbContract;

        public PositionDataService(FeelMeContext dbContract)
        {
            _dbContract = dbContract;
        }
        public virtual async Task<List<Position>> GetAllPositionAsync()
        {
            return await _dbContract.Positions.ToListAsync();
        }

        public virtual async Task<Position> GetPositionByIdAsync(int id)
        {
           return await _dbContract.Positions.FirstOrDefaultAsync(position => position.PositionId == id);
        }

        public virtual async Task InsertPositionAsync(Position position)
        {
            _dbContract.Add(position);
            await _dbContract.SaveChangesAsync();
        }

        public virtual async Task InsertListPositionAsync(List<Position> position)
        {
           _dbContract.AddRange(position);
            await _dbContract.SaveChangesAsync();
        }

        public virtual async Task UpdateListPositionAsync(List<Position> position)
        {
            _dbContract.UpdateRange(position);
            await _dbContract.SaveChangesAsync();
        }

        public virtual async Task UpdatePositionsAsync(Position position)
        {
            _dbContract.Update(position);
            await _dbContract.SaveChangesAsync();
        }
    }
}