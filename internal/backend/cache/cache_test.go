package cache

import (
	"testing"
	"time"
)

// getCache returns a new cache with the fake data
func getCache() (*Cache, error) {
	cache, err := New("../../test/data/cache.json")
	if err != nil {
		return nil, err
	}

	err = cache.Load()
	if err != nil {
		return nil, err
	}

	return cache, nil
}

// TestCacheLoadNoFile if we get an error then there's no cache file
func TestCacheLoadNoFile(t *testing.T) {
	// Create a cache with no file
	cache, err := New("../test/data/no-file")
	if err != nil {
		t.Fatalf("couldn't get default path: %v", err)
	}

	if err = cache.Load(); err != nil {
		t.Fatal("expected error, got nil")
	}
}

// TestCacheLoadCorrectly if we get an error then the cache file is bad
func TestCacheLoadCorrectly(t *testing.T) {
	// Create the cache object with a valid file
	cache, err := getCache()
	if err != nil {
		t.Fatalf("couldn't load the cache %v", err)
	}

	// Check if the cache is loaded correctly
	if len(cache.Content) != 1 {
		t.Fatal("expected 1 item in cache")
	}

	if _, ok := cache.Content["https://primordialsoup.info/feed"]; !ok {
		t.Fatal("expected https://primordialsoup.info/feed in cache")
	}
}

// TestCacheGetArticles if we get an error when there's a cache miss but the cache doesn't change
func TestCacheGetArticles(t *testing.T) {
	// Create the cache object with a valid file
	cache, err := getCache()
	if err != nil {
		t.Fatalf("couldn't load the cache %v", err)
	}

	// Check if the cache hit works
	_, err = cache.GetArticles("https://primordialsoup.info/feed")
	if err != nil {
		t.Fatalf("couldn't get article: %v", err)
	}

	if len(cache.Content) != 1 {
		t.Fatal("expected 1 item in cache")
	}

	// Check if the cache miss retrieves the item and puts it inside the cache
	_, err = cache.GetArticles("https://christitus.com/categories/virtualization/index.xml")
	if err != nil {
		t.Fatalf("couldn't get article: %v", err)
	}

	if len(cache.Content) != 2 {
		t.Fatal("expected 2 items in cache")
	}

	if _, ok := cache.Content["https://christitus.com/categories/virtualization/index.xml"]; !ok {
		t.Fatal("expected https://christitus.com/categories/virtualization/index.xml in cache")
	}
}

// TestCacheGetArticleExpired if we get an error then the store doesn't delete expired cache when getting data
func TestCacheGetArticleExpired(t *testing.T) {
	// Create the cache object with a valid file
	cache, err := getCache()
	if err != nil {
		t.Fatalf("couldn't load the cache %v", err)
	}

	// Get the item from the cache
	oldItem, ok := cache.Content["https://primordialsoup.info/feed"]
	if !ok {
		t.Fatal("expected https://primordialsoup.info/feed in cache")
	}

	// Make the item expired and insert it back into the map
	oldItem.Expire = time.Now().Add(-2 * DefaultCacheDuration)
	cache.Content["https://primordialsoup.info/feed"] = oldItem

	_, err = cache.GetArticles("https://primordialsoup.info/feed")
	if err != nil {
		t.Fatalf("couldn't get article: %v", err)
	}

	// Check if item expiry is updated (cache miss)
	newItem, ok := cache.Content["https://primordialsoup.info/feed"]
	if !ok {
		t.Fatal("expected https://primordialsoup.info/feed in cache")
	}

	if newItem.Expire.Equal(oldItem.Expire) {
		t.Fatal("expected the data to be refreshed and the expire to be updated")
	}
}
