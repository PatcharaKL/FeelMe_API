using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.CommentDataService
{
    public interface ICommentDataService
    {
         Task<List<Comment>> GetAllCommentAsync();
         Task<Comment> GetCommentByIdAsync(int id);
         Task InsertCommentAsync(Comment comment);
         Task InsertCommentAsync(List<Comment> comment);
         Task UpdateCommentsAsync(Comment comment);
         Task UpdateCommentAsync(List<Comment> comment);
    }
}