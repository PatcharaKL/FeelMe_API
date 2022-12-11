using Project_FeelMe.Data;
using Project_FeelMe.Models;
using Microsoft.EntityFrameworkCore;

namespace dotnet.Data.DataSevices.BoardDataService
{
    public class BoardDataService : IBoardDataService
    {
        private readonly FeelMeContext _dbContract;

        public BoardDataService(FeelMeContext dbContract)
        {
            _dbContract = dbContract;
        }
        public async Task<List<Board>> GetAllBoardAsync()
        {
            return await _dbContract.Boards.ToListAsync();
        }

        public async Task<Board> GetBoardByIdAsync(int id)
        {
           return await _dbContract.Boards.FirstOrDefaultAsync(board => board.BoardId == id);
        }

        public async Task InsertBoardAsync(Board board)
        {
            _dbContract.Add(board);
            await _dbContract.SaveChangesAsync();
        }

        public async Task InsertBoardAsync(List<Board> board)
        {
           _dbContract.AddRange(board);
            await _dbContract.SaveChangesAsync();
        }

        public async Task UpdateBoardAsync(List<Board> board)
        {
            _dbContract.UpdateRange(board);
            await _dbContract.SaveChangesAsync();
        }

        public async Task UpdateBoardsAsync(Board board)
        {
            _dbContract.Update(board);
            await _dbContract.SaveChangesAsync();
        }
    }
}