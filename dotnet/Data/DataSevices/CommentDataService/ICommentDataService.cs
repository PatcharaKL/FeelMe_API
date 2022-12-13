using Project_FeelMe.Models;

namespace dotnet.Data.DataSevices.CommentDataService
{
    public interface ICommentDataService
    {
         Task<List<Comment>> GetAllCommentAsync();
         Task<Comment> GetCommentByIdAsync(int id);
         Task InsertCommentAsync(Comment comment);
         Task InsertListCommentAsync(List<Comment> comment);
         Task UpdateCommentsAsync(Comment comment);
         Task UpdateListCommentAsync(List<Comment> comment);
    }
}