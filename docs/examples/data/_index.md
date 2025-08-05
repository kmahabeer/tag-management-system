# Sample Data

Below is mock data illustrating **many-to-many**, **hierarchies**, and **primary/secondary purposes**.

## 1. Assets

```text
id | filename
---|------------------------------
1  | fargo_still_01.jpg
2  | pulpfiction_suitcase.jpg
3  | hateful_eight_table.jpg
4  | datsun_240z.jpg
5  | figure_study_001.jpg
```

## 2. Tags

```text
id | name
---|--------------------------
1  | Quentin Tarantino
2  | Joel Coen
3  | Ethan Coen
4  | William H. Macy
5  | Frances McDormand
6  | Steve Buscemi
7  | Samuel L. Jackson
8  | Kurt Russell
9  | Film Still
10 | Script
11 | Figure Study
12 | Pulp Fiction
13 | The Hateful Eight
14 | Director
15 | Writer
16 | Actor
17 | Genre
18 | Crime
19 | Dark Comedy
20 | Drama
21 | Western
22 | Datsun 240z
23 | Car
24 | Vehicle
25 | Subject
26 | Character
27 | Prop
28 | Costume
29 | Movie
```

## 3. tag_types

```text
id | name
---|---------
1  | person
2  | role
3  | genre
4  | purpose
5  | vehicle
6  | subject
7  | character
8  | prop
9  | costume
10 | movie
```

## 4. tag_tag_types

```text
tag_id | tag_type_id
-------|-------------
1 (QT)      | 1 (person)
1           | 2 (role)
2 (Joel)    | 1
3 (Ethan)   | 1
4 (WHM)     | 1
5 (FM)      | 1
6 (SB)      | 1
7 (SLJ)     | 1
8 (KR)      | 1
14 (Director) | 2
15 (Writer)   | 2
16 (Actor)    | 2
17 (Genre)    | 3
18 (Crime)    | 3
19 (Dark Comedy) | 3
20 (Drama)      | 3
21 (Western)    | 3
9  (Film Still)| 4
10 (Script)     | 4
11 (Figure Study)| 4
22 (Datsun 240z)| 5
23 (Car)        | 5
24 (Vehicle)    | 5
25 (Subject)    | 6
26 (Character)  | 7
27 (Prop)       | 8
28 (Costume)    | 9
12 (Pulp Fiction)       | 10
13 (The Hateful Eight)  | 10
29 (Movie)              | 10
```

## 5. tag_relationships

```text
parent_tag_id | child_tag_id | relationship
--------------|--------------|-------------
14 (Director) | 1 (QT)       | role-of
14            | 2 (Joel)     | role-of
14            | 3 (Ethan)    | role-of
15 (Writer)   | 1            | role-of
15            | 3            | role-of
16 (Actor)    | 4            | role-of
16            | 5            | role-of
16            | 6            | role-of
16            | 7            | role-of
16            | 8            | role-of
12 (Pulp Fiction) | 1        | movie-of
12                  | 7      | movie-of
12                  | 6      | movie-of
12                  | 4      | movie-of
12                  | 5      | movie-of
12                  | 19     | movie-of
12                  | 18     | movie-of
12                  | 20     | movie-of
13 (The Hateful Eight) | 1    | movie-of
13                     | 8    | movie-of
13                     | 7    | movie-of
13                     | 16   | movie-of
13                     | 21   | movie-of
29 (Movie)           | 12      | category-of
29                   | 13      | category-of
```

## 6. asset_tags

```text
asset_id | tag_id
---------|-------
1        | 2 (Joel Coen)
1        | 3 (Ethan Coen)
1        | 4 (William H. Macy)
1        | 5 (Frances McDormand)
1        | 6 (Steve Buscemi)
1        | 7 (Samuel L. Jackson)
1        | 9 (Film Still)
1        | 12 (Pulp Fiction)
1        | 26 (Character) -- e.g. multiple instances in child table for each character
1        | 28 (Costume)      -- e.g. "Black Suit"
1        | 27 (Prop)         -- e.g. "Gun"

2        | 1 (Quentin Tarantino)
2        | 7 (Samuel L. Jackson)
2        | 9 (Film Still)
2        | 12 (Pulp Fiction)

3        | 7 (Samuel L. Jackson)
3        | 8 (Kurt Russell)
3        | 9 (Film Still)
3        | 13 (The Hateful Eight)

4        | 22 (Datsun 240z)
4        | 23 (Car)
4        | 24 (Vehicle)

5        | 25 (Subject)     -- e.g., "Model"
5        | 11 (Figure Study)
5        | 26 (Character)? -- or different subject tags
```

## 7. asset_purposes

```text
asset_id | purpose_tag_id | is_primary
---------|----------------|-----------
1        | 9 (Film Still)   | true
1        | 29 (Movie reference) | false
2        | 9                | true
2        | 29               | false
3        | 9                | true
4        | 11 (Figure Study)| true
5        | 11               | true
```
