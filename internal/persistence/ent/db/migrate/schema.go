// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PackageColumns holds the columns for the "package" table.
	PackageColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Size: 2048},
		{Name: "workflow_id", Type: field.TypeString, Size: 255},
		{Name: "run_id", Type: field.TypeUUID, Unique: true},
		{Name: "aip_id", Type: field.TypeUUID, Nullable: true},
		{Name: "location_id", Type: field.TypeUUID, Nullable: true},
		{Name: "status", Type: field.TypeInt8},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "started_at", Type: field.TypeTime, Nullable: true},
		{Name: "completed_at", Type: field.TypeTime, Nullable: true},
	}
	// PackageTable holds the schema information for the "package" table.
	PackageTable = &schema.Table{
		Name:       "package",
		Columns:    PackageColumns,
		PrimaryKey: []*schema.Column{PackageColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "package_name_idx",
				Unique:  false,
				Columns: []*schema.Column{PackageColumns[1]},
				Annotation: &entsql.IndexAnnotation{
					Prefix: 50,
				},
			},
			{
				Name:    "package_aip_id_idx",
				Unique:  false,
				Columns: []*schema.Column{PackageColumns[4]},
			},
			{
				Name:    "package_location_id_idx",
				Unique:  false,
				Columns: []*schema.Column{PackageColumns[5]},
			},
			{
				Name:    "package_status_idx",
				Unique:  false,
				Columns: []*schema.Column{PackageColumns[6]},
			},
			{
				Name:    "package_created_at_idx",
				Unique:  false,
				Columns: []*schema.Column{PackageColumns[7]},
			},
			{
				Name:    "package_started_at_idx",
				Unique:  false,
				Columns: []*schema.Column{PackageColumns[8]},
			},
		},
	}
	// PreservationActionColumns holds the columns for the "preservation_action" table.
	PreservationActionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "workflow_id", Type: field.TypeString, Size: 255},
		{Name: "type", Type: field.TypeInt8},
		{Name: "status", Type: field.TypeInt8},
		{Name: "started_at", Type: field.TypeTime, Nullable: true},
		{Name: "completed_at", Type: field.TypeTime, Nullable: true},
		{Name: "package_id", Type: field.TypeInt},
	}
	// PreservationActionTable holds the schema information for the "preservation_action" table.
	PreservationActionTable = &schema.Table{
		Name:       "preservation_action",
		Columns:    PreservationActionColumns,
		PrimaryKey: []*schema.Column{PreservationActionColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "preservation_action_package_preservation_actions",
				Columns:    []*schema.Column{PreservationActionColumns[6]},
				RefColumns: []*schema.Column{PackageColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// PreservationTaskColumns holds the columns for the "preservation_task" table.
	PreservationTaskColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "task_id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Size: 2048},
		{Name: "status", Type: field.TypeInt8},
		{Name: "started_at", Type: field.TypeTime, Nullable: true},
		{Name: "completed_at", Type: field.TypeTime, Nullable: true},
		{Name: "note", Type: field.TypeString, Size: 2147483647},
		{Name: "preservation_action_id", Type: field.TypeInt},
	}
	// PreservationTaskTable holds the schema information for the "preservation_task" table.
	PreservationTaskTable = &schema.Table{
		Name:       "preservation_task",
		Columns:    PreservationTaskColumns,
		PrimaryKey: []*schema.Column{PreservationTaskColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "preservation_task_preservation_action_tasks",
				Columns:    []*schema.Column{PreservationTaskColumns[7]},
				RefColumns: []*schema.Column{PreservationActionColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PackageTable,
		PreservationActionTable,
		PreservationTaskTable,
	}
)

func init() {
	PackageTable.Annotation = &entsql.Annotation{
		Table: "package",
	}
	PreservationActionTable.ForeignKeys[0].RefTable = PackageTable
	PreservationActionTable.Annotation = &entsql.Annotation{
		Table: "preservation_action",
	}
	PreservationTaskTable.ForeignKeys[0].RefTable = PreservationActionTable
	PreservationTaskTable.Annotation = &entsql.Annotation{
		Table: "preservation_task",
	}
}
