import { ref, onMounted } from 'vue'
import axios from 'axios'

export function usePostJobToolbar() {
    const toolbarData = ref([])
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/postjob/toolbar')
            toolbarData.value = res.data
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { toolbarData, loading, error }
}
