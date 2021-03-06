package repository

import (
  "context"
  "log"
  "cloud.google.com/firestore"
  "google.golang.org/api/iterator"
  "go-mux-api/entity"
)

const (
  projectId = ""
)

type repo struct { }

func NewPostRepository() PostReposiory {
  return &repo{}
}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
  // Firestore database
  ctx := context.Background()
  client, err := firestore.NewClient(ctx, projectId)
  if err != nil {
    log.Fatalf("Failed to create a firestore client: %v", err)
    return nil, err
  }

  defer client.Close()

  _, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
    "ID": post.Id,
    "Title": post.Title,
    "Text": post.Text,
  })

  if err != nil {
    log.Fatalf("Failed adding a new post: %v", err)
    return nil, err
  }

  return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
  // Firestore database
  ctx := context.Background()
  client, err := firestore.NewClient(ctx, projectId)
  if err != nil {
    log.Fatalf("Failed to create a firestore client: %v", err)
    return nil, err
  }

  defer client.Close()
  var post []entity.Post
  iterator := client.Collection(collectionName).Documents(ctx)

  for {
    doc, err := iterator.Next()
    if err != nil {
      log.Fatalf("Failed to iterate the list of posts: %v", err)
      return nil, err
    }
    post := entity.Post{
      Id: doc.Data()["Id"].(int),
      Title: doc.Data()["Title"].(string),
      Text: doc.Data()["Text"].(string),
    }

    posts = append(posts, post)
  }

  return posts, nil
}
