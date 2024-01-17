<template>
    <div ref="editorContainer" class="editor-container"></div>
</template>

<script lang="ts">
import { defineComponent, onMounted, ref, watch } from 'vue';
import * as monaco from 'monaco-editor';
import yaml from 'js-yaml';

export default defineComponent({
    name: 'MonacoEditor',
    props: {
        content: {
            type: String,
            required: true,
        },
    },
    setup(props) {
        const editorContainer = ref<HTMLElement | null>(null);
        let editor: monaco.editor.IStandaloneCodeEditor;
        const currentContent = ref(props.content);

        function validateYaml(content: string) {
            try {
                yaml.load(content);
                return [];
            } catch (e) {
                if (e instanceof Error) {
                    return [
                        {
                            startLineNumber: 1,
                            startColumn: 1,
                            endLineNumber: 1,
                            endColumn: 1,
                            message: e.message,
                            severity: monaco.MarkerSeverity.Error,
                        },
                    ];
                }
                // generic error
                console.error('Unexpected error validating YAML file:', e);
                return [
                    {
                        startLineNumber: 1,
                        startColumn: 1,
                        endLineNumber: 1,
                        endColumn: 1,
                        message: 'Unexpected error validating YAML file',
                        severity: monaco.MarkerSeverity.Error,
                    },
                ];
            }
        }

        function getCurrentContent() {
            return currentContent.value;
        }

        onMounted(() => {
            editor = monaco.editor.create(editorContainer.value!, {
                value: props.content,
                language: 'yaml',
                theme: 'vs-dark',
            });
            // validate
            const markers = validateYaml(props.content);
            monaco.editor.setModelMarkers(editor.getModel()!, 'yaml', markers);

            // listen changes
            editor.onDidChangeModelContent(() => {
                const model = editor.getModel();
                const content = model?.getValue() || '';
                const markers = validateYaml(content);
                monaco.editor.setModelMarkers(model!, 'yaml', markers);
                currentContent.value = model?.getValue() || '';
            });
        });

        watch(
            () => props.content,
            (newValue) => {
                if (editor) {
                    editor.setValue(newValue);
                    const markers = validateYaml(newValue);
                    monaco.editor.setModelMarkers(editor.getModel()!, 'yaml', markers);
                }
            },
        );

        return { editorContainer, currentContent, getCurrentContent };
    },
});
</script>

<style scoped>
.editor-container {
    height: 500px;
}
</style>
