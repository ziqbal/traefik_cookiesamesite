# traefik_cookiesamesite



# cookiesamesite = plugin GO module/src
# plugin_cookesamesite = local name for plugin

# ROUTER
```
            - --experimental.localPlugins.plugin_cookesamesite.modulename=cookiesamesite
```

# plugin_cookesamesite = local name for plugin
# mw_cookiesamesite = local name for middleware

# WEB APP LABELS

```
            - "traefik.http.routers.$omaad_SERVICE.middlewares=mw_cookiesamesite"
            #- "traefik.http.middlewares.mw_cookiesamesite.plugin.plugin_cookesamesite.rewrites.header=Set-Cookie"
            #- "traefik.http.middlewares.mwcookiesamesite.plugin.plugin_cookesamesite.rewrites.regex=^(.*)$$"
            - "traefik.http.middlewares.mw_cookiesamesite.plugin.plugin_cookesamesite.rewrites.replacement=None"

```
