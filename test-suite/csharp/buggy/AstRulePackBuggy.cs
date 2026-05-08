using System.Collections.Generic;
using System.Threading.Tasks;

public static class AstRulePackBuggy
{
    public static async Task RunAsync(object gate, IEnumerable<int> values)
    {
        Task.Run(() => 42);
        _ = Task.Factory.StartNew(() => 7);
        lock (gate)
        {
            await Task.Delay(1);
        }
        Parallel.ForEach(values, async item => { await Task.Delay(item); });
    }
}
