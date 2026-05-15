import { ref, onMounted } from 'vue'
import axios from 'axios'

export function useFooter() {
    const footer = ref(null)
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/home/footer')
            footer.value = res.data
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { footer, loading, error }
}
