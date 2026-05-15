import { ref, onMounted } from 'vue'
import axios from 'axios'

export function usePostJobPreview() {
    const previewData = ref(null)
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/postjob/preview')
            previewData.value = res.data
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { previewData, loading, error }
}
