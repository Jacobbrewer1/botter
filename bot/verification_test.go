package bot

import (
	"testing"
)

func TestVerifyMultiRole(t *testing.T) {
	tests := []struct {
		name        string
		memberRoles []string
		roles       []role
		level       int
		expected    bool
	}{
		{"Does not have Member permissions", TestingSetupMemberRoles(0), roles, member.level, false},
		{"Has Member permissions", TestingSetupMemberRoles(member.level), roles, member.level, true},
		{"Does not have Mod permissions", TestingSetupMemberRoles(squad.level), roles, moderator.level, false},
		{"Has Mod permissions", TestingSetupMemberRoles(moderator.level), roles, moderator.level, true},
		{"Has Squad permissions", TestingSetupMemberRoles(squad.level), roles, squad.level, true},
		{"Does not have Squad permissions", TestingSetupMemberRoles(member.level), roles, squad.level, false},
		{"Does not have Boss permissions", TestingSetupMemberRoles(moderator.level), roles, boss.level, false},
		{"Has Boss permissions", TestingSetupMemberRoles(moderator.level), roles, boss.level, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool := verifyMultiRole(tt.memberRoles, tt.level)
			if gotBool != tt.expected {
				t.Errorf("verifyMultiRole() = %v, expected %v", gotBool, tt.expected)
			}
		})
	}
}

func TestVerifySingleRole(t *testing.T) {
	tests := []struct {
		name         string
		memberRoles  []string
		roleRequired role
		expected     bool
	}{
		{"Does not have Mod permissions", TestingSetupMemberRoles(member.level), moderator, false},
		{"Has Mod permissions", TestingSetupMemberRoles(moderator.level), moderator, true},
		{"Has Squad permissions", TestingSetupMemberRoles(squad.level), squad, true},
		{"Does not have Boss", TestingSetupMemberRoles(moderator.level), boss, false},
		{"Does not have Member", TestingSetupMemberRoles(0), member, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBool := verifySingleRole(tt.memberRoles, tt.roleRequired.id)
			if gotBool != tt.expected {
				t.Errorf("verifySingleRole() = %v, expected %v", gotBool, tt.expected)
			}
		})
	}
}
