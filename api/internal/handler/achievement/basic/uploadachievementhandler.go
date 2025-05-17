package basic

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/bearllfleed/scholar-track/pkg/response"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/bearllfleed/scholar-track/api/internal/logic/achievement/basic"
	"github.com/bearllfleed/scholar-track/api/internal/svc"
	"github.com/bearllfleed/scholar-track/api/internal/types"
	//"github.com/zeromicro/go-zero/rest/httpx" // No longer needed for parsing request body
)

func UploadAchievementHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Parse Multipart Form Data (including files and form fields)
		// 32 << 20 limits the *total* size of the request body stored in memory
		// (excluding file contents stored on disk if they exceed memory limits)
		err := r.ParseMultipartForm(32 << 20) // 32MB max memory
		if err != nil {
			logx.Errorf("Failed to parse multipart form: %v", err)
			// Consider returning a more specific error, e.g., http.StatusRequestEntityTooLarge
			response.ErrWithMessage(r.Context(), w, "请求解析失败或大小超过限制")
			return
		}

		// 2. Get File Headers from the parsed form
		filesMap := r.MultipartForm.File // map[string][]*multipart.FileHeader

		// 3. Validate required files
		materials := filesMap["materials"]
		if len(materials) == 0 {
			response.ErrWithMessage(r.Context(), w, "请上传证明材料")
			return
		}
		// Note: We are NOT reading file content here in the handler.
		// The logic layer will handle opening and processing the files using the headers.

		// 4. Manually parse basic information fields from form values
		var req types.UploadAchievementReq

		req.Code = r.FormValue("code")
		req.Name = r.FormValue("name")
		req.Description = r.FormValue("description")

		// --- Parse non-string fields with error handling ---

		// Level (int32)
		levelStr := r.FormValue("level")
		if levelStr == "" {
			response.ErrWithMessage(r.Context(), w, "level 参数不能为空")
			return
		}
		levelInt, err := strconv.ParseInt(levelStr, 10, 32)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "level 参数格式错误")
			return
		}
		req.Level = int32(levelInt)

		// Rank (int32)
		rankStr := r.FormValue("rank")
		if rankStr == "" {
			response.ErrWithMessage(r.Context(), w, "rank 参数不能为空")
			return
		}
		rankInt, err := strconv.ParseInt(rankStr, 10, 32)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "rank 参数格式错误")
			return
		}
		req.Rank = int32(rankInt)

		// AwardTime (int64)
		awardTimeStr := r.FormValue("awardTime")
		if awardTimeStr == "" {
			response.ErrWithMessage(r.Context(), w, "awardTime 参数不能为空")
			return
		}
		awardTimeInt, err := strconv.ParseInt(awardTimeStr, 10, 64)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "awardTime 参数格式错误")
			return
		}
		req.AwardTime = awardTimeInt

		// Share (bool)
		shareStr := r.FormValue("share")
		if shareStr == "" {
			response.ErrWithMessage(r.Context(), w, "share 参数不能为空")
			return
		}
		req.Share, err = strconv.ParseBool(shareStr)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "share 参数格式错误 (应为 true 或 false)")
			return
		}

		// Star (uint64, optional)
		starStr := r.FormValue("star")
		if starStr != "" {
			starUint, err := strconv.ParseUint(starStr, 10, 64)
			if err != nil {
				response.ErrWithMessage(r.Context(), w, "star 参数格式错误")
				return
			}
			req.Star = starUint
		}

		claims, ok := r.Context().Value("claims").(*types.CustomClaims)
		if ok {
			req.Uploader = claims.BaseClaims.ID
		} else {
			response.ErrWithMessage(r.Context(), w, "无法认证用户信息，请重新登录")
			return
		}

		// Team (bool)
		teamStr := r.FormValue("team")
		if teamStr == "" {
			response.ErrWithMessage(r.Context(), w, "team 参数不能为空")
			return
		}
		req.Team, err = strconv.ParseBool(teamStr)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "team 参数格式错误 (应为 true 或 false)")
			return
		}

		// TeamMembers ([]uint64) - Assuming comma-separated string "1,2,3"
		teamMembersStr := r.FormValue("teamMembers")
		if req.Team && teamMembersStr == "" {
			// If it's a team achievement, members might be required depending on logic
			response.ErrWithMessage(r.Context(), w, "团队成果必须提供 teamMembers 参数")
			return
		}
		if teamMembersStr != "" {
			members := strings.Split(teamMembersStr, ",")
			req.TeamMembers = make([]uint64, 0, len(members))
			for _, memberStr := range members {
				memberUint, err := strconv.ParseUint(strings.TrimSpace(memberStr), 10, 64)
				if err != nil {
					response.ErrWithMessage(r.Context(), w, "teamMembers 参数格式错误 (应为逗号分隔的数字)")
					return
				}
				req.TeamMembers = append(req.TeamMembers, memberUint)
			}
		}

		// CategoryId (uint64)
		categoryIdStr := r.FormValue("categoryId")
		if categoryIdStr == "" {
			response.ErrWithMessage(r.Context(), w, "categoryId 参数不能为空")
			return
		}
		categoryIdUint, err := strconv.ParseUint(categoryIdStr, 10, 64)
		if err != nil {
			response.ErrWithMessage(r.Context(), w, "categoryId 参数格式错误")
			return
		}
		req.CategoryId = categoryIdUint

		// OtherInfo (map[string]interface{}) - Assuming JSON string in form field
		otherInfoStr := r.FormValue("otherInfo")
		if otherInfoStr != "" {
			err := json.Unmarshal([]byte(otherInfoStr), &req.OtherInfo)
			if err != nil {
				response.ErrWithMessage(r.Context(), w, "otherInfo 参数格式错误 (应为有效的 JSON 字符串)")
				return
			}
		}

		// Status (int32, optional, default=1)
		statusStr := r.FormValue("status")
		if statusStr != "" {
			statusInt, err := strconv.ParseInt(statusStr, 10, 32)
			if err != nil {
				response.ErrWithMessage(r.Context(), w, "status 参数格式错误")
				return
			}
			req.Status = int32(statusInt)
		} else {
			req.Status = 1 // Default value if not provided
		}

		// 5. Call the logic layer, passing the parsed request data and the file headers map
		l := basic.NewUploadAchievementLogic(r.Context(), svcCtx)
		// IMPORTANT: Assumes the logic layer's UploadAchievement method signature is updated to:
		// func (l *UploadAchievementLogic) UploadAchievement(req *types.UploadAchievementReq, files map[string][]*multipart.FileHeader) (*types.UploadAchievementResp, error)
		_, err = l.UploadAchievement(&req, filesMap) // Pass parsed req and file headers
		if err != nil {
			// Provide more context about the error if possible
			logx.Errorf("UploadAchievement logic error: %v", err)
			// Consider mapping specific logic errors to user-friendly messages
			response.ErrWithMessage(r.Context(), w, "上传失败: "+err.Error()) // Show specific error for now
		} else {
			response.Success(r.Context(), w) // Return success response with data (e.g., new achievement ID)
		}
	}
}
