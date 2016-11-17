
type
  OrganizationTreeItem* = object
    children*: seq[OrganizationTreeItem]
    globalid*: string
