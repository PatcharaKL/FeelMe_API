using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.PositionDataService
{
    public interface IPositionDataService
    {
        Task<List<Position>> GetAllPositionAsync();
         Task<Position> GetPositionByIdAsync(int id);
         Task InsertPositionAsync(Position position);
         Task InsertPositionAsync(List<Position> position);
         Task UpdatePositionsAsync(Position position);
         Task UpdatePositionAsync(List<Position> position);
    }
}