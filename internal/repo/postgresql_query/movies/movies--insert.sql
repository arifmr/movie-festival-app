INSERT INTO movies (
    title,
    description,
    duration,
    artists,
    genres,
    url
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
) Returning id;