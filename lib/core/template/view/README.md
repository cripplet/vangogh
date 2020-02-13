Vangogh Core Views
----

These are contextual templates and are manually added to `template.ParseFiles`,
as opposed to the `component/` directory which may be globbed together and
are guaranteed to have no name collisions.

The user is responsible for manually ensuring each page has a fully defined
template when rendering.

View templates should ideally be kept to a single file (per view) to make
loading easier.

TODO(minkezhang): Define the template name which these views should define.
