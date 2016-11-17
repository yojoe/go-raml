
import EnumJoinOrganizationInvitationRole
import times
type
  JoinOrganizationInvitation* = object
    created*: Time
    organization*: string
    role*: EnumJoinOrganizationInvitationRole
    user*: string
