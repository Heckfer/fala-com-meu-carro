package handler
import (
	"github.com/drborges/appx"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/heckfer/fala-com-meu-carro/model"
	"net/http"
	"fmt"
	"github.com/heckfer/fala-com-meu-carro/resources"
)

func FlagPostHandler(r render.Render, params martini.Params, appx *appx.Datastore) {
	postId := params["post_id"]

	response := model.Response{
		ErrorCode: http.StatusOK,
		Message: []string{},
		Data: nil,
	}

	post := model.Post{}
	post.SetEncodedKey(postId)
	err := appx.Load(&post)

	fmt.Print("aqui")
	if err != nil {
		fmt.Print("not found")
		fmt.Print(err)
		response.Message = []string{err.Error()}
		response.ErrorCode = http.StatusBadRequest
	} else {
		fmt.Print("found")
		fmt.Print(err)
		post.Flagged = true
		err = appx.Save(&post)

		if err != nil {
			fmt.Print("save problem")
			fmt.Print(err)
			response.Message = []string{err.Error()}
			response.ErrorCode = http.StatusInternalServerError
		} else {
			postResource := &resources.PostResource{}
			postResource.From(post)
			response.Data = postResource
		}
	}

	r.JSON(200, response)
}