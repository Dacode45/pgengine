package pgengine

//TriggerType map[TriggerName]map[EventName]ActionName
type TriggerType map[string]map[string]string

type Trigger map[string]Action

NewTrigger(trigger Trigger) Trigger {
	if _, ok := trigger["OnEnter"]; !ok {
		trigger["OnEnter"] = EmptyAction
	}
	if _, ok := trigger["OnExit"]; !ok {
		trigger["OnExit"] = EmptyAction
	}
	if _, ok := trigger["OnUse"]; !ok {
		trigger["OnUse"] = EmptyAction
	}
	return trigger 
}
