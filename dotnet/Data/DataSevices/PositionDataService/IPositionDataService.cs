using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.PositionDataService
{
    public interface IPositionDataService
    {
        Task<List<Position>> GetAllPositionAsync();
         Task<Position> GetPositionByIdAsync(int id);
         Task InsertPositionAsync(Position position);
         Task InsertListPositionAsync(List<Position> position);
         Task UpdatePositionsAsync(Position position);
         Task UpdateListPositionAsync(List<Position> position);
    }
}