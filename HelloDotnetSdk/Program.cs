using GitHub;
using GitHub.Client;
using GitHub.Authentication;

var token = Environment.GetEnvironmentVariable("GITHUB_TOKEN") ?? "";
var request = RequestAdapter.Create(new TokenAuthenticationProvider("Octokit.Gen", token));
var gitHubClient = new GitHubClient(request);

var pullRequests = await gitHubClient.Repos["octokit"]["octokit.net"].Pulls.GetAsync();

foreach (var pullRequest in pullRequests)
{
    Console.WriteLine($"#{pullRequest.Number} {pullRequest.Title}");
}
