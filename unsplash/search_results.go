// Copyright (c) 2017 Hardik Bagdi <hbagdi1@binghamton.edu>
//
// MIT License
//
// Permission is hereby granted, free of charge, to any person obtaining
// a copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to
// permit persons to whom the Software is furnished to do so, subject to
// the following conditions:
//
// The above copyright notice and this permission notice shall be
// included in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
// EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
// NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
// LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
// WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package unsplash

// UserSearchResult represents the result for a search for users.
type UserSearchResult struct {
	Total      *int    `json:"total"`
	TotalPages *int    `json:"total_pages"`
	Results    *[]User `json:"results"`
}

// PhotoSearchResult represents the result for a search for photos.
type PhotoSearchResult struct {
	Total      *int     `json:"total"`
	TotalPages *int     `json:"total_pages"`
	Results    *[]Photo `json:"results"`
}

// CollectionSearchResult represents the result for a search for collections.
type CollectionSearchResult struct {
	Total      *int          `json:"total"`
	TotalPages *int          `json:"total_pages"`
	Results    *[]Collection `json:"results"`
}
