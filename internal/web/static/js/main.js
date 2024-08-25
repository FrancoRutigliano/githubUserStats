function userName(user) {
    fetch("user-activity/?user=" + user)
        .then(Response => Response.json())
}