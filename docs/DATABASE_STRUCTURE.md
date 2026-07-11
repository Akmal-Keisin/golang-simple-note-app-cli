# Table Structure

## `notes`

| Column name  | Type (inferred)                   | Notes                               |
| ------------ | --------------------------------- | ----------------------------------- |
| `id`         | integer / serial / identity       | Primary key used to identify a note |
| `title`      | text or varchar                   | Note title                          |
| `content`    | text                              | Note body/content                   |
| `created_at` | timestamp / timestamptz           | Set when a note is created          |
| `updated_at` | timestamp / timestamptz, nullable | Set when a note is updated          |

```sql
CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NULL
);
```
