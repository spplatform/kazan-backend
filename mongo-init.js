db.createUser(
    {
        user: "kazan",
        pwd: "kazan",
        roles: [
            {
                role: "readWrite",
                db: "kazandb"
            }
        ]
    }
);