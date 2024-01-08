# traefik_cookiesamesite
This Traefik plugin overcame an issue with browsers requiring cookies to have the cookie samesite header set correctly.

My specific issue was related to OpenAM and Google Chrome integration.

If the back end application sets the samesite in the cookie header then this plugin lets it through otherwise it injects the samesite key value into the cookie header.

The value to be injected can be set using the container service labels.

I'll make this readme more useful if anyone is truly interested.


```
cookiesamesite = plugin GO module/src
plugin_cookesamesite = local name for plugin

ROUTER
            - --experimental.localPlugins.plugin_cookesamesite.modulename=cookiesamesite

plugin_cookesamesite = local name for plugin
mw_cookiesamesite = local name for middleware

WEB APP LABELS

            - "traefik.http.routers.$omaad_SERVICE.middlewares=mw_cookiesamesite"
            #- "traefik.http.middlewares.mw_cookiesamesite.plugin.plugin_cookesamesite.rewrites.header=Set-Cookie"
            #- "traefik.http.middlewares.mwcookiesamesite.plugin.plugin_cookesamesite.rewrites.regex=^(.*)$$"
            - "traefik.http.middlewares.mw_cookiesamesite.plugin.plugin_cookesamesite.rewrites.replacement=None"

```
