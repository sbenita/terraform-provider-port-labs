package webhook

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func WebhookSchema() map[string]schema.Attribute {
	return map[string]schema.Attribute{
		"id": schema.StringAttribute{
			Computed: true,
		},
		"identifier": schema.StringAttribute{
			MarkdownDescription: "The identifier of the webhook",
			Optional:            true,
			Computed:            true,
		},
		"title": schema.StringAttribute{
			MarkdownDescription: "The title of the webhook",
			Optional:            true,
		},
		"icon": schema.StringAttribute{
			MarkdownDescription: "The icon of the webhook",
			Optional:            true,
		},
		"description": schema.StringAttribute{
			MarkdownDescription: "The description of the webhook",
			Optional:            true,
		},
		"enabled": schema.BoolAttribute{
			MarkdownDescription: "Whether the webhook is enabled",
			Optional:            true,
			Computed:            true,
			Default:             booldefault.StaticBool(false),
		},
		"security": schema.SingleNestedAttribute{
			MarkdownDescription: "The security of the webhook",
			Optional:            true,
			Attributes: map[string]schema.Attribute{
				"secret": schema.StringAttribute{
					MarkdownDescription: "The secret of the webhook",
					Optional:            true,
				},
				"signature_header_name": schema.StringAttribute{
					MarkdownDescription: "The signature header name of the webhook",
					Optional:            true,
				},
				"signature_algorithm": schema.StringAttribute{
					MarkdownDescription: "The signature algorithm of the webhook",
					Optional:            true,
				},
				"signature_prefix": schema.StringAttribute{
					MarkdownDescription: "The signature prefix of the webhook",
					Optional:            true,
				},
				"request_identifier_path": schema.StringAttribute{
					MarkdownDescription: "The request identifier path of the webhook",
					Optional:            true,
				},
			},
		},
		"mappings": schema.ListNestedAttribute{
			MarkdownDescription: "The mappings of the webhook",
			Optional:            true,
			NestedObject: schema.NestedAttributeObject{
				Attributes: map[string]schema.Attribute{
					"blueprint": schema.StringAttribute{
						MarkdownDescription: "The blueprint of the mapping",
						Required:            true,
					},
					"filter": schema.StringAttribute{
						MarkdownDescription: "The filter of the mapping",
						Optional:            true,
					},
					"items_to_parse": schema.StringAttribute{
						MarkdownDescription: "The items to parser of the mapping",
						Optional:            true,
					},
					"entity": schema.SingleNestedAttribute{
						MarkdownDescription: "The entity of the mapping",
						Required:            true,
						Attributes: map[string]schema.Attribute{
							"identifier": schema.StringAttribute{
								MarkdownDescription: "The identifier of the entity",
								Required:            true,
							},
							"title": schema.StringAttribute{
								MarkdownDescription: "The title of the entity",
								Optional:            true,
							},
							"icon": schema.StringAttribute{
								MarkdownDescription: "The icon of the entity",
								Optional:            true,
							},
							"team": schema.StringAttribute{
								MarkdownDescription: "The team of the entity",
								Optional:            true,
							},
							"properties": schema.MapAttribute{
								MarkdownDescription: "The properties of the entity",
								Optional:            true,
								ElementType:         types.StringType,
							},
							"relations": schema.MapAttribute{
								MarkdownDescription: "The relations of the entity",
								Optional:            true,
								ElementType:         types.StringType,
							},
						},
					},
				},
			},
		},
		"created_at": schema.StringAttribute{
			MarkdownDescription: "The creation date of the webhook",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"created_by": schema.StringAttribute{
			MarkdownDescription: "The creator of the webhook",
			Computed:            true,
			PlanModifiers: []planmodifier.String{
				stringplanmodifier.UseStateForUnknown(),
			},
		},
		"updated_at": schema.StringAttribute{
			MarkdownDescription: "The last update date of the webhook",
			Computed:            true,
		},
		"updated_by": schema.StringAttribute{
			MarkdownDescription: "The last updater of the webhook",
			Computed:            true,
		},
	}
}

func (r *WebhookResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "Webhook resource",
		Attributes:          WebhookSchema(),
	}
}
