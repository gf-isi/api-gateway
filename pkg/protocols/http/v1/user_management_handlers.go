package v1

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/influenzanet/api-gateway/pkg/utils"
	"google.golang.org/grpc/status"

	umAPI "github.com/influenzanet/user-management-service/pkg/api"
)

func (h *HttpEndpoints) loginWithEmailAsParticipantHandl(c *gin.Context) {
	var req umAPI.LoginWithEmailMsg
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.AsParticipant = true

	token, err := h.clients.UserManagement.LoginWithEmail(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, token)
}

func (h *HttpEndpoints) resendVerificationCodeHandl(c *gin.Context) {
	var req umAPI.SendVerificationCodeReq
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.clients.UserManagement.SendVerificationCode(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) getVerificationCodeWithTokenHandl(c *gin.Context) {
	var req umAPI.AutoValidateReq
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.clients.UserManagement.AutoValidateTempToken(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) loginWithEmailForManagementHandl(c *gin.Context) {
	var req umAPI.LoginWithEmailMsg
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.AsParticipant = false

	resp, err := h.clients.UserManagement.LoginWithEmail(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) signupWithEmailHandl(c *gin.Context) {
	var req umAPI.SignupWithEmailMsg
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.clients.UserManagement.SignupWithEmail(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) switchProfileHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)
	var req umAPI.SwitchProfileRequest
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Token = token
	resp, err := h.clients.UserManagement.SwitchProfile(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) tokenRenewHandl(c *gin.Context) {
	var req umAPI.RefreshJWTRequest
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.AccessToken = c.MustGet("encodedToken").(string)
	token, err := h.clients.UserManagement.RenewJWT(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, token)
}

func (h *HttpEndpoints) userPasswordChangeHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	var req umAPI.PasswordChangeMsg
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Token = token

	resp, err := h.clients.UserManagement.ChangePassword(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		log.Println(st.Message())
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) changeAccountEmailHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	var req umAPI.EmailChangeMsg
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Token = token
	resp, err := h.clients.UserManagement.ChangeAccountIDEmail(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) userSetPreferredLanguageHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	var req umAPI.LanguageChangeMsg
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Token = token
	resp, err := h.clients.UserManagement.ChangePreferredLanguage(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) getUserHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	userID := c.Param("id")

	userRefReq := &umAPI.UserReference{
		Token:  token,
		UserId: userID,
	}

	resp, err := h.clients.UserManagement.GetUser(context.Background(), userRefReq)
	if err != nil {
		st := status.Convert(err)
		log.Println(st.Message())
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) deleteAccountHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	var req umAPI.UserReference
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Token = token
	resp, err := h.clients.UserManagement.DeleteAccount(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		log.Println(st.Message())
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) saveProfileHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	var req umAPI.ProfileRequest
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Token = token
	resp, err := h.clients.UserManagement.SaveProfile(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) removeProfileHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	var req umAPI.ProfileRequest
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Token = token
	resp, err := h.clients.UserManagement.RemoveProfile(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) userUpdateContactPreferencesHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	var req umAPI.ContactPreferencesMsg
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Token = token
	resp, err := h.clients.UserManagement.UpdateContactPreferences(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) userAddEmailHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	var req umAPI.ContactInfoMsg
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Token = token
	resp, err := h.clients.UserManagement.AddEmail(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) userRemoveEmailHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	var req umAPI.ContactInfoMsg
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Token = token
	resp, err := h.clients.UserManagement.RemoveEmail(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) createUserHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	var req umAPI.CreateUserReq
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Token = token
	resp, err := h.clients.UserManagement.CreateUser(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) findNonParticipantUsersHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	var req umAPI.FindNonParticipantUsersMsg
	req.Token = token
	resp, err := h.clients.UserManagement.FindNonParticipantUsers(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) userAddRoleHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	var req umAPI.RoleMsg
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Token = token
	resp, err := h.clients.UserManagement.AddRoleForUser(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) userRemoveRoleHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	var req umAPI.RoleMsg
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Token = token
	resp, err := h.clients.UserManagement.RemoveRoleForUser(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) revokeRefreshTokensHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)

	var req umAPI.RevokeRefreshTokensReq
	req.Token = token
	resp, err := h.clients.UserManagement.RevokeAllRefreshTokens(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) initiatePasswordResetHandl(c *gin.Context) {
	var req umAPI.InitiateResetPasswordMsg
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.clients.UserManagement.InitiatePasswordReset(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) getInfosForPasswordResetHandl(c *gin.Context) {
	var req umAPI.GetInfosForResetPasswordMsg
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.clients.UserManagement.GetInfosForPasswordReset(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) passwordResetHandl(c *gin.Context) {
	var req umAPI.ResetPasswordMsg
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.clients.UserManagement.ResetPassword(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) verifyUserContactHandl(c *gin.Context) {
	var req umAPI.TempToken
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp, err := h.clients.UserManagement.VerifyContact(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) resendContanctVerificationEmailHandl(c *gin.Context) {
	token := c.MustGet("validatedToken").(*umAPI.TokenInfos)
	var req umAPI.ResendContactVerificationReq
	if err := h.JsonToProto(c, &req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.Token = token
	resp, err := h.clients.UserManagement.ResendContactVerification(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}

func (h *HttpEndpoints) unsubscribeNewsletterHandl(c *gin.Context) {
	var req umAPI.TempToken
	req.Token = c.DefaultQuery("token", "")
	resp, err := h.clients.UserManagement.UseUnsubscribeToken(context.Background(), &req)
	if err != nil {
		st := status.Convert(err)
		c.JSON(utils.GRPCStatusToHTTP(st.Code()), gin.H{"error": st.Message()})
		return
	}
	h.SendProtoAsJSON(c, http.StatusOK, resp)
}