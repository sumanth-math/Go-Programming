using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Barclays.Theater.DataAccessLogic.Interface
{
    // Repository Interface
    public interface IRepository<T>
    {
        // Create Repository
        void Create(IEnumerable<T> entity);

        // Update repository
        void UpdateAll(IEnumerable<T> entit);

        // Get repository
        IEnumerable<T> GetAll();

        // Delete
        void Delete(T entity);
    }
}
