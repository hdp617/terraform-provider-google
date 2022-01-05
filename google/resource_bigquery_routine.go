// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/structure"
  "github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
  "github.com/hashicorp/terraform-plugin-sdk/v2/diag"
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/logging"
)




func resourceBigQueryRoutine() *schema.Resource {
    return &schema.Resource{
        Create: resourceBigQueryRoutineCreate,
        Read: resourceBigQueryRoutineRead,
        Update: resourceBigQueryRoutineUpdate,
        Delete: resourceBigQueryRoutineDelete,

        Importer: &schema.ResourceImporter{
            State: resourceBigQueryRoutineImport,
        },

        Timeouts: &schema.ResourceTimeout {
            Create: schema.DefaultTimeout(4 * time.Minute),
            Update: schema.DefaultTimeout(4 * time.Minute),
            Delete: schema.DefaultTimeout(4 * time.Minute),
        },



        Schema: map[string]*schema.Schema{
"definition_body": {
    Type: schema.TypeString,
    Required: true,
	Description: `The body of the routine. For functions, this is the expression in the AS clause.
If language=SQL, it is the substring inside (but excluding) the parentheses.`,
},
			"dataset_id": {
    Type: schema.TypeString,
    Required: true,
  ForceNew: true,
	Description: `The ID of the dataset containing this routine`,
},
			"routine_id": {
    Type: schema.TypeString,
    Required: true,
  ForceNew: true,
	Description: `The ID of the the routine. The ID must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_). The maximum length is 256 characters.`,
},
	
"arguments": {
    Type: schema.TypeList,
    Optional: true,
	Description: `Input/output argument of a function or a stored procedure.`,
                Elem: &schema.Resource{
        Schema: map[string]*schema.Schema{
                      "argument_kind": {
    Type: schema.TypeString,
    Optional: true,
	ValidateFunc: validation.StringInSlice([]string{"FIXED_TYPE","ANY_TYPE",""}, false),
	Description: `Defaults to FIXED_TYPE. Default value: "FIXED_TYPE" Possible values: ["FIXED_TYPE", "ANY_TYPE"]`,
    Default: "FIXED_TYPE",
},
                      "data_type": {
    Type: schema.TypeString,
    Optional: true,
		ValidateFunc: validation.StringIsJSON,
		StateFunc: func(v interface{}) string { s, _ := structure.NormalizeJsonString(v); return s },
	Description: `A JSON schema for the data type. Required unless argumentKind = ANY_TYPE.
~>**NOTE**: Because this field expects a JSON string, any changes to the string
will create a diff, even if the JSON itself hasn't changed. If the API returns
a different value for the same schema, e.g. it switched the order of values
or replaced STRUCT field type with RECORD field type, we currently cannot
suppress the recurring diff this causes. As a workaround, we recommend using
the schema as returned by the API.`,
},
                      "mode": {
    Type: schema.TypeString,
    Optional: true,
	ValidateFunc: validation.StringInSlice([]string{"IN","OUT","INOUT",""}, false),
	Description: `Specifies whether the argument is input or output. Can be set for procedures only. Possible values: ["IN", "OUT", "INOUT"]`,
},
                      "name": {
    Type: schema.TypeString,
    Optional: true,
	Description: `The name of this argument. Can be absent for function return argument.`,
},
                  },
      },
        },
"description": {
    Type: schema.TypeString,
    Optional: true,
	Description: `The description of the routine if defined.`,
},
"determinism_level": {
    Type: schema.TypeString,
    Optional: true,
	ValidateFunc: validation.StringInSlice([]string{"DETERMINISM_LEVEL_UNSPECIFIED","DETERMINISTIC","NOT_DETERMINISTIC",""}, false),
	Description: `The determinism level of the JavaScript UDF if defined. Possible values: ["DETERMINISM_LEVEL_UNSPECIFIED", "DETERMINISTIC", "NOT_DETERMINISTIC"]`,
},
"imported_libraries": {
    Type: schema.TypeList,
    Optional: true,
	Description: `Optional. If language = "JAVASCRIPT", this field stores the path of the
imported JAVASCRIPT libraries.`,
            Elem: &schema.Schema{
        Type: schema.TypeString,
      },
    },
"language": {
    Type: schema.TypeString,
    Optional: true,
	ValidateFunc: validation.StringInSlice([]string{"SQL","JAVASCRIPT",""}, false),
	Description: `The language of the routine. Possible values: ["SQL", "JAVASCRIPT"]`,
},
"return_table_type": {
    Type: schema.TypeString,
    Optional: true,
		ValidateFunc: validation.StringIsJSON,
		StateFunc: func(v interface{}) string { s, _ := structure.NormalizeJsonString(v); return s },
	Description: `Optional. Can be set only if routineType = "TABLE_VALUED_FUNCTION".

If absent, the return table type is inferred from definitionBody at query time in each query
that references this routine. If present, then the columns in the evaluated table result will
be cast to match the column types specificed in return table type, at query time.`,
},
"return_type": {
    Type: schema.TypeString,
    Optional: true,
		ValidateFunc: validation.StringIsJSON,
		StateFunc: func(v interface{}) string { s, _ := structure.NormalizeJsonString(v); return s },
	Description: `A JSON schema for the return type. Optional if language = "SQL"; required otherwise.
If absent, the return type is inferred from definitionBody at query time in each query
that references this routine. If present, then the evaluated result will be cast to
the specified returned type at query time. ~>**NOTE**: Because this field expects a JSON
string, any changes to the string will create a diff, even if the JSON itself hasn't
changed. If the API returns a different value for the same schema, e.g. it switche
d the order of values or replaced STRUCT field type with RECORD field type, we currently
cannot suppress the recurring diff this causes. As a workaround, we recommend using
the schema as returned by the API.`,
},
"routine_type": {
    Type: schema.TypeString,
    Optional: true,
  ForceNew: true,
	ValidateFunc: validation.StringInSlice([]string{"SCALAR_FUNCTION","PROCEDURE","TABLE_VALUED_FUNCTION",""}, false),
	Description: `The type of routine. Possible values: ["SCALAR_FUNCTION", "PROCEDURE", "TABLE_VALUED_FUNCTION"]`,
},
"creation_time": {
    Type: schema.TypeInt,
    Computed: true,
	Description: `The time when this routine was created, in milliseconds since the
epoch.`,
},
"last_modified_time": {
    Type: schema.TypeInt,
    Computed: true,
	Description: `The time when this routine was modified, in milliseconds since the
epoch.`,
},
            "project": {
                Type:     schema.TypeString,
                Optional: true,
                Computed: true,
                ForceNew: true,
            },
        },
        UseJSONNumber: true,
    }
}



func resourceBigQueryRoutineCreate(d *schema.ResourceData, meta interface{}) error {
    config := meta.(*Config)
    userAgent, err := generateUserAgentString(d, config.userAgent)
    if err != nil {
        return err
    }

    obj := make(map[string]interface{})
        routineReferenceProp, err := expandBigQueryRoutineRoutineReference(nil, d, config)
    if err != nil {
        return err
    } else if !isEmptyValue(reflect.ValueOf(routineReferenceProp)) {
        obj["routineReference"] = routineReferenceProp
    }
        routineTypeProp, err := expandBigQueryRoutineRoutineType(d.Get( "routine_type" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("routine_type"); !isEmptyValue(reflect.ValueOf(routineTypeProp)) && (ok || !reflect.DeepEqual(v, routineTypeProp)) {
        obj["routineType"] = routineTypeProp
    }
        languageProp, err := expandBigQueryRoutineLanguage(d.Get( "language" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("language"); !isEmptyValue(reflect.ValueOf(languageProp)) && (ok || !reflect.DeepEqual(v, languageProp)) {
        obj["language"] = languageProp
    }
        argumentsProp, err := expandBigQueryRoutineArguments(d.Get( "arguments" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("arguments"); !isEmptyValue(reflect.ValueOf(argumentsProp)) && (ok || !reflect.DeepEqual(v, argumentsProp)) {
        obj["arguments"] = argumentsProp
    }
        returnTypeProp, err := expandBigQueryRoutineReturnType(d.Get( "return_type" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("return_type"); !isEmptyValue(reflect.ValueOf(returnTypeProp)) && (ok || !reflect.DeepEqual(v, returnTypeProp)) {
        obj["returnType"] = returnTypeProp
    }
        returnTableTypeProp, err := expandBigQueryRoutineReturnTableType(d.Get( "return_table_type" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("return_table_type"); !isEmptyValue(reflect.ValueOf(returnTableTypeProp)) && (ok || !reflect.DeepEqual(v, returnTableTypeProp)) {
        obj["returnTableType"] = returnTableTypeProp
    }
        importedLibrariesProp, err := expandBigQueryRoutineImportedLibraries(d.Get( "imported_libraries" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("imported_libraries"); !isEmptyValue(reflect.ValueOf(importedLibrariesProp)) && (ok || !reflect.DeepEqual(v, importedLibrariesProp)) {
        obj["importedLibraries"] = importedLibrariesProp
    }
        definitionBodyProp, err := expandBigQueryRoutineDefinitionBody(d.Get( "definition_body" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("definition_body"); !isEmptyValue(reflect.ValueOf(definitionBodyProp)) && (ok || !reflect.DeepEqual(v, definitionBodyProp)) {
        obj["definitionBody"] = definitionBodyProp
    }
        descriptionProp, err := expandBigQueryRoutineDescription(d.Get( "description" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
        obj["description"] = descriptionProp
    }
        determinismLevelProp, err := expandBigQueryRoutineDeterminismLevel(d.Get( "determinism_level" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("determinism_level"); !isEmptyValue(reflect.ValueOf(determinismLevelProp)) && (ok || !reflect.DeepEqual(v, determinismLevelProp)) {
        obj["determinismLevel"] = determinismLevelProp
    }



    url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}/routines")
    if err != nil {
        return err
    }

    log.Printf("[DEBUG] Creating new Routine: %#v", obj)
    billingProject := ""

    project, err := getProject(d, config)
    if err != nil {
        return fmt.Errorf("Error fetching project for Routine: %s", err)
    }
    billingProject = project


    // err == nil indicates that the billing_project value was found
    if bp, err := getBillingProject(d, config); err == nil {
      billingProject = bp
    }

    res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
    if err != nil {
        return fmt.Errorf("Error creating Routine: %s", err)
    }
                                                                                                                                                    
    // Store the ID now
    id, err := replaceVars(d, config, "projects/{{project}}/datasets/{{dataset_id}}/routines/{{routine_id}}")
    if err != nil {
        return fmt.Errorf("Error constructing id: %s", err)
    }
    d.SetId(id)




    log.Printf("[DEBUG] Finished creating Routine %q: %#v", d.Id(), res)

    return resourceBigQueryRoutineRead(d, meta)
}


func resourceBigQueryRoutineRead(d *schema.ResourceData, meta interface{}) error {
    config := meta.(*Config)
    userAgent, err := generateUserAgentString(d, config.userAgent)
    if err != nil {
        return err
    }

    url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}/routines/{{routine_id}}")
    if err != nil {
        return err
    }

    billingProject := ""

    project, err := getProject(d, config)
    if err != nil {
        return fmt.Errorf("Error fetching project for Routine: %s", err)
    }
    billingProject = project


    // err == nil indicates that the billing_project value was found
    if bp, err := getBillingProject(d, config); err == nil {
      billingProject = bp
    }

    res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
    if err != nil {
        return handleNotFoundError(err, d, fmt.Sprintf("BigQueryRoutine %q", d.Id()))
    }


    if err := d.Set("project", project); err != nil {
        return fmt.Errorf("Error reading Routine: %s", err)
    }


// Terraform must set the top level schema field, but since this object contains collapsed properties
// it's difficult to know what the top level should be. Instead we just loop over the map returned from flatten.
    if flattenedProp := flattenBigQueryRoutineRoutineReference(res["routineReference"], d, config); flattenedProp != nil {
        if gerr, ok := flattenedProp.(*googleapi.Error); ok {
			return fmt.Errorf("Error reading Routine: %s", gerr)
		}
        casted := flattenedProp.([]interface{})[0]
        if casted != nil {
            for k, v := range casted.(map[string]interface{}) {
                if err := d.Set(k, v); err != nil {
                	return fmt.Errorf("Error setting %s: %s", k, err)
                }
            }
        }
    }
    if err := d.Set("routine_type", flattenBigQueryRoutineRoutineType(res["routineType"], d, config)); err != nil {
        return fmt.Errorf("Error reading Routine: %s", err)
    }
    if err := d.Set("creation_time", flattenBigQueryRoutineCreationTime(res["creationTime"], d, config)); err != nil {
        return fmt.Errorf("Error reading Routine: %s", err)
    }
    if err := d.Set("last_modified_time", flattenBigQueryRoutineLastModifiedTime(res["lastModifiedTime"], d, config)); err != nil {
        return fmt.Errorf("Error reading Routine: %s", err)
    }
    if err := d.Set("language", flattenBigQueryRoutineLanguage(res["language"], d, config)); err != nil {
        return fmt.Errorf("Error reading Routine: %s", err)
    }
    if err := d.Set("arguments", flattenBigQueryRoutineArguments(res["arguments"], d, config)); err != nil {
        return fmt.Errorf("Error reading Routine: %s", err)
    }
    if err := d.Set("return_type", flattenBigQueryRoutineReturnType(res["returnType"], d, config)); err != nil {
        return fmt.Errorf("Error reading Routine: %s", err)
    }
    if err := d.Set("return_table_type", flattenBigQueryRoutineReturnTableType(res["returnTableType"], d, config)); err != nil {
        return fmt.Errorf("Error reading Routine: %s", err)
    }
    if err := d.Set("imported_libraries", flattenBigQueryRoutineImportedLibraries(res["importedLibraries"], d, config)); err != nil {
        return fmt.Errorf("Error reading Routine: %s", err)
    }
    if err := d.Set("definition_body", flattenBigQueryRoutineDefinitionBody(res["definitionBody"], d, config)); err != nil {
        return fmt.Errorf("Error reading Routine: %s", err)
    }
    if err := d.Set("description", flattenBigQueryRoutineDescription(res["description"], d, config)); err != nil {
        return fmt.Errorf("Error reading Routine: %s", err)
    }
    if err := d.Set("determinism_level", flattenBigQueryRoutineDeterminismLevel(res["determinismLevel"], d, config)); err != nil {
        return fmt.Errorf("Error reading Routine: %s", err)
    }

    return nil
}

func resourceBigQueryRoutineUpdate(d *schema.ResourceData, meta interface{}) error {
    config := meta.(*Config)
    userAgent, err := generateUserAgentString(d, config.userAgent)
    if err != nil {
    	return err
    }

    billingProject := ""

    project, err := getProject(d, config)
    if err != nil {
        return fmt.Errorf("Error fetching project for Routine: %s", err)
    }
    billingProject = project


    obj := make(map[string]interface{})
            routineReferenceProp, err := expandBigQueryRoutineRoutineReference(nil, d, config)
    if err != nil {
        return err
    } else if !isEmptyValue(reflect.ValueOf(routineReferenceProp)) {
        obj["routineReference"] = routineReferenceProp
    }
            routineTypeProp, err := expandBigQueryRoutineRoutineType(d.Get( "routine_type" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("routine_type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, routineTypeProp)) {
        obj["routineType"] = routineTypeProp
    }
            languageProp, err := expandBigQueryRoutineLanguage(d.Get( "language" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("language"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, languageProp)) {
        obj["language"] = languageProp
    }
            argumentsProp, err := expandBigQueryRoutineArguments(d.Get( "arguments" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("arguments"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, argumentsProp)) {
        obj["arguments"] = argumentsProp
    }
            returnTypeProp, err := expandBigQueryRoutineReturnType(d.Get( "return_type" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("return_type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, returnTypeProp)) {
        obj["returnType"] = returnTypeProp
    }
            returnTableTypeProp, err := expandBigQueryRoutineReturnTableType(d.Get( "return_table_type" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("return_table_type"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, returnTableTypeProp)) {
        obj["returnTableType"] = returnTableTypeProp
    }
            importedLibrariesProp, err := expandBigQueryRoutineImportedLibraries(d.Get( "imported_libraries" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("imported_libraries"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, importedLibrariesProp)) {
        obj["importedLibraries"] = importedLibrariesProp
    }
            definitionBodyProp, err := expandBigQueryRoutineDefinitionBody(d.Get( "definition_body" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("definition_body"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, definitionBodyProp)) {
        obj["definitionBody"] = definitionBodyProp
    }
            descriptionProp, err := expandBigQueryRoutineDescription(d.Get( "description" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
        obj["description"] = descriptionProp
    }
            determinismLevelProp, err := expandBigQueryRoutineDeterminismLevel(d.Get( "determinism_level" ), d, config)
    if err != nil {
        return err
    } else if v, ok := d.GetOkExists("determinism_level"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, determinismLevelProp)) {
        obj["determinismLevel"] = determinismLevelProp
    }



    url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}/routines/{{routine_id}}")
    if err != nil {
        return err
    }

    log.Printf("[DEBUG] Updating Routine %q: %#v", d.Id(), obj)

    // err == nil indicates that the billing_project value was found
    if bp, err := getBillingProject(d, config); err == nil {
      billingProject = bp
    }

    res, err := sendRequestWithTimeout(config, "PUT", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

    if err != nil {
        return fmt.Errorf("Error updating Routine %q: %s", d.Id(), err)
    } else {
	log.Printf("[DEBUG] Finished updating Routine %q: %#v", d.Id(), res)
    }


    return resourceBigQueryRoutineRead(d, meta)
}

func resourceBigQueryRoutineDelete(d *schema.ResourceData, meta interface{}) error {
    config := meta.(*Config)
    userAgent, err := generateUserAgentString(d, config.userAgent)
    if err != nil {
    	return err
    }


    billingProject := ""

    project, err := getProject(d, config)
    if err != nil {
        return fmt.Errorf("Error fetching project for Routine: %s", err)
    }
    billingProject = project


    url, err := replaceVars(d, config, "{{BigQueryBasePath}}projects/{{project}}/datasets/{{dataset_id}}/routines/{{routine_id}}")
    if err != nil {
        return err
    }

    var obj map[string]interface{}
    log.Printf("[DEBUG] Deleting Routine %q", d.Id())

    // err == nil indicates that the billing_project value was found
    if bp, err := getBillingProject(d, config); err == nil {
      billingProject = bp
    }

    res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
    if err != nil {
        return handleNotFoundError(err, d, "Routine")
    }


    log.Printf("[DEBUG] Finished deleting Routine %q: %#v", d.Id(), res)
    return nil
}

func resourceBigQueryRoutineImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
    config := meta.(*Config)
    if err := parseImportId([]string{
        "projects/(?P<project>[^/]+)/datasets/(?P<dataset_id>[^/]+)/routines/(?P<routine_id>[^/]+)",
        "(?P<project>[^/]+)/(?P<dataset_id>[^/]+)/(?P<routine_id>[^/]+)",
        "(?P<dataset_id>[^/]+)/(?P<routine_id>[^/]+)",
    }, d, config); err != nil {
      return nil, err
    }

    // Replace import id for the resource id
    id, err := replaceVars(d, config, "projects/{{project}}/datasets/{{dataset_id}}/routines/{{routine_id}}")
    if err != nil {
        return nil, fmt.Errorf("Error constructing id: %s", err)
    }
    d.SetId(id)


    return []*schema.ResourceData{d}, nil
}

func flattenBigQueryRoutineRoutineReference(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  if v == nil {
    return nil
  }
  original := v.(map[string]interface{})
    if len(original) == 0 {
    return nil
  }
    transformed := make(map[string]interface{})
          transformed["dataset_id"] =
    flattenBigQueryRoutineRoutineReferenceDatasetId(original["datasetId"], d, config)
              transformed["routine_id"] =
    flattenBigQueryRoutineRoutineReferenceRoutineId(original["routineId"], d, config)
        return []interface{}{transformed}
}
      func flattenBigQueryRoutineRoutineReferenceDatasetId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

      func flattenBigQueryRoutineRoutineReferenceRoutineId(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

  

func flattenBigQueryRoutineRoutineType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

func flattenBigQueryRoutineCreationTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenBigQueryRoutineLastModifiedTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenBigQueryRoutineLanguage(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

func flattenBigQueryRoutineArguments(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  if v == nil {
    return v
  }
  l := v.([]interface{})
  transformed := make([]interface{}, 0, len(l))
  for _, raw := range l {
    original := raw.(map[string]interface{})
    if len(original) < 1 {
      // Do not include empty json objects coming back from the api
      continue
    }
    transformed = append(transformed, map[string]interface{}{
          "name": flattenBigQueryRoutineArgumentsName(original["name"], d, config),
          "argument_kind": flattenBigQueryRoutineArgumentsArgumentKind(original["argumentKind"], d, config),
          "mode": flattenBigQueryRoutineArgumentsMode(original["mode"], d, config),
          "data_type": flattenBigQueryRoutineArgumentsDataType(original["dataType"], d, config),
        })
  }
  return transformed
}
      func flattenBigQueryRoutineArgumentsName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

      func flattenBigQueryRoutineArgumentsArgumentKind(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

      func flattenBigQueryRoutineArgumentsMode(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

      func flattenBigQueryRoutineArgumentsDataType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	b, err := json.Marshal(v)
	if err != nil {
        // TODO: return error once https://github.com/GoogleCloudPlatform/magic-modules/issues/3257 is fixed.
		log.Printf("[ERROR] failed to marshal schema to JSON: %v", err)
	}
	return string(b)
}

  

func flattenBigQueryRoutineReturnType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	b, err := json.Marshal(v)
	if err != nil {
        // TODO: return error once https://github.com/GoogleCloudPlatform/magic-modules/issues/3257 is fixed.
		log.Printf("[ERROR] failed to marshal schema to JSON: %v", err)
	}
	return string(b)
}

func flattenBigQueryRoutineReturnTableType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	b, err := json.Marshal(v)
	if err != nil {
        // TODO: return error once https://github.com/GoogleCloudPlatform/magic-modules/issues/3257 is fixed.
		log.Printf("[ERROR] failed to marshal schema to JSON: %v", err)
	}
	return string(b)
}

func flattenBigQueryRoutineImportedLibraries(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

func flattenBigQueryRoutineDefinitionBody(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

func flattenBigQueryRoutineDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}

func flattenBigQueryRoutineDeterminismLevel(v interface{}, d *schema.ResourceData, config *Config) interface{} {
  return v
}


func expandBigQueryRoutineRoutineReference(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {

        transformed := make(map[string]interface{})
        transformed["datasetId"] = d.Get("dataset_id")
        project, _ := getProject(d, config)
        transformed["projectId"] = project
        transformed["routineId"] = d.Get("routine_id")

        return transformed, nil
}



func expandBigQueryRoutineRoutineType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
  return v, nil
}



func expandBigQueryRoutineLanguage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
  return v, nil
}



func expandBigQueryRoutineArguments(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
  l := v.([]interface{})
  req := make([]interface{}, 0, len(l))
  for _, raw := range l {
    if raw == nil {
      continue
    }
    original := raw.(map[string]interface{})
    transformed := make(map[string]interface{})

      transformedName, err := expandBigQueryRoutineArgumentsName(original["name"], d, config)
      if err != nil {
        return nil, err
      } else if val := reflect.ValueOf(transformedName); val.IsValid() && !isEmptyValue(val) {
        transformed["name"] = transformedName      }

      transformedArgumentKind, err := expandBigQueryRoutineArgumentsArgumentKind(original["argument_kind"], d, config)
      if err != nil {
        return nil, err
      } else if val := reflect.ValueOf(transformedArgumentKind); val.IsValid() && !isEmptyValue(val) {
        transformed["argumentKind"] = transformedArgumentKind      }

      transformedMode, err := expandBigQueryRoutineArgumentsMode(original["mode"], d, config)
      if err != nil {
        return nil, err
      } else if val := reflect.ValueOf(transformedMode); val.IsValid() && !isEmptyValue(val) {
        transformed["mode"] = transformedMode      }

      transformedDataType, err := expandBigQueryRoutineArgumentsDataType(original["data_type"], d, config)
      if err != nil {
        return nil, err
      } else if val := reflect.ValueOf(transformedDataType); val.IsValid() && !isEmptyValue(val) {
        transformed["dataType"] = transformedDataType      }

    req = append(req, transformed)
  }
  return req, nil
}






func expandBigQueryRoutineArgumentsName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
  return v, nil
}




func expandBigQueryRoutineArgumentsArgumentKind(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
  return v, nil
}




func expandBigQueryRoutineArgumentsMode(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
  return v, nil
}


func expandBigQueryRoutineArgumentsDataType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func expandBigQueryRoutineReturnType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}

func expandBigQueryRoutineReturnTableType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	b := []byte(v.(string))
	if len(b) == 0 {
		return nil, nil
	}
	m := make(map[string]interface{})
	if err := json.Unmarshal(b, &m); err != nil {
		return nil, err
	}
	return m, nil
}



func expandBigQueryRoutineImportedLibraries(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
  return v, nil
}



func expandBigQueryRoutineDefinitionBody(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
  return v, nil
}



func expandBigQueryRoutineDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
  return v, nil
}



func expandBigQueryRoutineDeterminismLevel(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
  return v, nil
}
