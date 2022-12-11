using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.BoardDataService
{
    public interface IBoardDataService
    {
         Task<List<Board>> GetAllBoardAsync();
         Task<Board> GetBoardByIdAsync(int id);
         Task InsertBoardAsync(Board board);
         Task InsertBoardAsync(List<Board> board);
         Task UpdateBoardsAsync(Board board);
         Task UpdateBoardAsync(List<Board> board);
    }
}