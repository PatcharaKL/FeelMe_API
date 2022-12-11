using Microsoft.EntityFrameworkCore;
using Project_FeelMe.Data;
using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.CommentDataService
{
    public class CommentDataService : ICommentDataService
    {
        private readonly FeelMeContext _dbContract;
        public CommentDataService(FeelMeContext dbContract)
        {
            _dbContract = dbContract;
        }
        public virtual async Task<List<Comment>> GetAllCommentAsync()
        {
            return await _dbContract.Comments.ToListAsync();
        }

        public virtual async Task<Comment> GetCommentByIdAsync(int id)
        {
            return await _dbContract.Comments.FirstOrDefaultAsync(comment => comment.CommentId == id);
        }

        public virtual async Task InsertCommentAsync(Comment comment)
        {
             _dbContract.Add(comment);
            await _dbContract.SaveChangesAsync();
        }

        public virtual async Task InsertCommentAsync(List<Comment> comment)
        {
           _dbContract.AddRange(comment);
            await _dbContract.SaveChangesAsync();
        }

        public virtual async Task UpdateCommentAsync(List<Comment> comment)
        {
            _dbContract.UpdateRange(comment);
            await _dbContract.SaveChangesAsync();
        }

        public virtual async Task UpdateCommentsAsync(Comment comment)
        {
             _dbContract.Update(comment);
            await _dbContract.SaveChangesAsync();
        }
    }
}