| command     | description                                                         |            status |
| :---------- | :------------------------------------------------------------------ | ----------------: |
| up          | Connect to Tailscale, logging in if needed                          |   won't implement |
| down        | Disconnect from Tailscale                                           | may not implement |
| set         | Change specified preferences                                        |   not yet started |
| login       | Log in to a Tailscale account                                       |   won't implement |
| logout      | Disconnect from Tailscale and expire current node key               |   won't implement |
| switch      | Switch to a different Tailscale account                             |   won't implement |
| configure   | Configure the host to enable more Tailscale features                |   won't implement |
| syspolicy   | Diagnose the MDM and system policy configuration                    |   won't implement |
| netcheck    | Print an analysis of local network conditions                       | may not implement |
| ip          | Show Tailscale IP addresses                                         |              todo |
| dns         | Diagnose the internal DNS forwarder                                 |              todo |
| status      | Show state of tailscaled and its connections                        |       in progress |
| metrics     | Show Tailscale metrics                                              |              todo |
| ping        | Ping a host at the Tailscale layer, see how it routed               |              todo |
| nc          | Connect to a port on a host, connected to stdin/stdout              | may not implement |
| ssh         | SSH to a Tailscale machine                                          | may not implement |
| funnel      | Serve content and local servers on the internet                     |              todo |
| serve       | Serve content and local servers on your tailnet                     |              todo |
| version     | Print Tailscale version                                             |              todo |
| web         | Run a web server for controlling Tailscale                          |              todo |
| file        | Send or receive files                                               |              todo |
| bugreport   | Print a shareable identifier to help diagnose issues                | may not implement |
| cert        | Get TLS certs                                                       | may not implement |
| lock        | Manage tailnet lock                                                 |              todo |
| licenses    | Get open source license information                                 |   won't implement |
| exit-node   | Show machines on your tailnet configured as exit nodes              | may not implement |
| update      | Update Tailscale to the latest/different version                    |   won't implement |
| whois       | Show the machine and user associated with a Tailscale IP (v4 or v6) |              todo |
| appc-routes | Print the current app connector routes                              |              todo |
| completion  | Shell tab-completion scripts                                        |   won't implement |

FLAGS
--socket value
path to tailscaled socket (default /var/run/tailscaled.socket)

## tested version

1.94.1 tailscale commit: 62c6f1cd7560763edcc552251abadf7dd4659f82 long version: 1.94.1-t62c6f1cd7-g09fea6572 other commit: 09fea6572e9c7a839b4a8c209420141fdb33fbb4 go version: go1.25.5
